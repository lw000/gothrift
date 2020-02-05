namespace go echo

struct EchoRequest {
	1: string msg;
	2: i32  tag;
}

struct EchoResponse {
	1: string msg;
	2: i32 tag
}

service Echo {
	EchoResponse echo(1: EchoRequest req);
}