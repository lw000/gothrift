namespace go tapi

struct ReqRegist {
	1: string account;
	2: string password;
}

struct AckRegist {
	1: i32 code;
	2: string message;
	3: string account;
}

struct ReqLogin {
	1: string account;
	2: string password;
}

struct AckLogin {
	1: i32 code;
	2: string message;
}

service TapiService {
		AckLogin login(1: ReqLogin req);
    	AckRegist regist(1: ReqRegist req);
}
