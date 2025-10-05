package ch14l3

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	timeCost    uint32 = 3
	memoryCost  uint32 = 32 * 1024
	parallelism uint8  = 4
	keyLen             = 32
	saltLen            = 16
)

// argon2 패스워드 해쉬 함수
func hashPassword(password string) (string, error) {
	// ?

	// salt 생성
	salt := make([]byte, saltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// KDF argon2를 이용해 password 해쉬
	hashed := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, parallelism, keyLen)

	// base64로 salt와 hashed 인코딩
	b64 := base64.RawStdEncoding
	encodedSalt := b64.EncodeToString(salt)
	encodedHash := b64.EncodeToString(hashed)

	// phc formatted string 반환
	return fmt.Sprintf("$argon2id$v=19$m=%s,t=%s,p=%s$%s$%s", strconv.FormatUint(uint64(memoryCost), 10), strconv.FormatUint(uint64(timeCost), 10), strconv.FormatUint(uint64(parallelism), 10), encodedSalt, encodedHash), nil
	// @@@ strconv.FormatUint(변환할 uint64, 진법 n)은 주어진 uint64를 base n 진법으로 표현한 것을 string으로 반환
}

// argon2 해쉬 검증 함수
func checkPasswordHash(password, hash string) bool {
	// ?

	// hash는 $argon2id$v=19$m=<memKiB>,t=<time>,p=<parallelism>$<saltBase64>$<hashBase64> 형태
	// Expect: ["", "argon2id", "v=19", "m=..,t=..,p=..", "<saltB64>", "<hashB64>"]
	parts := strings.Split(hash, "$")
	if len(parts) != 6 || parts[1] != "argon2id" || parts[2] != "v=19" {
		return false
	}

	// "m=..,t=..,p=.."를 ,로 분리
	params := strings.Split(parts[3], ",")

	// @@@ strconv.ParseUint(s, base, bitsize)는 자열을 부호 없는 정수(uint64)로 변환
	// @@@ @@@ 변환할 입력 문자열 s가 무슨 진법으로 표현되어 있는지 base 인자로 제공해야 한다
	// @@@ @@@ 결과로 반환할 uint의 비트 수는 원본 const 비트 수와 동일하게 입력
	m, err := strconv.ParseUint(params[0][2:], 10, 32)
	if err != nil {
		return false
	}
	t, err := strconv.ParseUint(params[1][2:], 10, 32)
	if err != nil {
		return false
	}
	p, err := strconv.ParseUint(params[2][2:], 10, 8)
	if err != nil {
		return false
	}

	// base64 인코딩된 salt와 hash를 디코딩해 []byte로 변환
	b64 := base64.RawStdEncoding
	salt, err := b64.DecodeString(parts[4])
	if err != nil {
		return false
	}
	want, err := b64.DecodeString(parts[5])
	if err != nil {
		return false
	}

	// KDF argon2를 이용해 새롭게 password 해쉬해 기존의 해쉬값(want)와 비교해서 검증
	got := argon2.IDKey([]byte(password), salt, uint32(t), uint32(m), uint8(p), uint32(len(want)))
	// Avoid timing leaks:
	// @@@ 단순 비교하면 앞자리부터 차례로 비교하다가 처음으로 불일치하는 순간 바로 비교가 종료되기 때문에
	// @@@ 공격자가 그 비교 위치에 따라 반환 시간이 달라지는 것을 이용해 일부 내용 추측이 가능함
	// @@@ ex: 공격자가 3개의 블록을 제대로 알고 4번째 블록을 brute force 하는 경우
	// @@@ => 틀린 것을 입력할때는 4번쨰 블록 비교에서 종료되고 맞는 것을 입력할 때는 5번째 블록 비교에서 종료하므로 시간이 달라진다
	// @@@ @@@ subtle.ConstantTimeCompare는 두 슬라이스의 길이가 다르면 바로 0을 반환하고
	// @@@ @@@ 길이가 같으면 반드시 슬라이스의 끝까지 전부 비교한 후 결과를 반환하므로
	// @@@ @@@ 실행 시간은 입력 값의 내용과 무관하게(상수 시간, constant time) 슬라이스 길이에만 비례
	return subtle.ConstantTimeCompare(got, want) == 1
}
