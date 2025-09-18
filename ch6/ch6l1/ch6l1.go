package ch6l1

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
	// <-chan 는 read only channel, chan<-는 write only channel
	// ?

	// @@@ defer 사용해서 채널 close 하기
	defer close(result)
	// defer를 사용하면 for 루프 안에 break 대신 return을 사용해도 문제 없다

	for {
		t, tok := <-textCh
		// 채널이 닫히고 채널 안에 데이터가 없으면 tok 가 false
		// false일 때 t는 해당 타입의 zero 값
		if !tok {
			// tok가 false ==> 채널이 close 상태이면서 동시에 남은 데이터도 없음
			// @@@ text와 key는 동일 길이 이므로 두 채널 각각의 마지막 글자가 나와서 XOR 연산이 될때까진
			// @@@ tok가 false가 되지 않는다 ===> 따라서 keyCh에 데이터가 남았는데 for 루프가 종료되는 일이 일어나지 않는다
			break
		}
		k, kok := <-keyCh
		if !kok {
			break
		}

		// if !tok || !kok {
		// 	// 두 채널 중 하나라도 데이터가 빈 close 채널이면 for 루프 종료
		// 	break
		// }

		result <- (t ^ k)
	}

	// close(result)
}
