package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "https://rycimagespixel.s3.us-east-1.amazonaws.com/578be853144861.5929aa1414698.gif?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=ASIA5G2VGWW5UWSQXRF2%2F20250810%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20250810T005643Z&X-Amz-Expires=300&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEJH%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJIMEYCIQDJOljOhEpAK%2BwMN9ak6gt7fYX7jI8B3MelFRnHMaWfHwIhALaChKJzHZeMeh%2FJcMhwgLgfuhqxIM94nvj%2BOnM7D1S7KuMCCMr%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEQABoMOTA4MDI3NDA5ODUxIgzeZ8hQOFIDX46QFkIqtwLsp89R%2FIQUP5v8zyqj8DCUpNWRyZ6v8evofR7kqwfSzxRz3wxuAxGqPH16GhgoH9ksXKIiH3DQQB08avXznxwevx8SqprmNgyDPKOgIC7F6DiOS%2B6MxGH4e%2FLTnkEw6NYngTaFNlBE38mAxlFg2XSUz2bFTdPLylci4u1%2F6xxnA3v3E2ixwL1Wm6Q%2BgUNJK5Wv%2BJ0VZ6R32L9SS%2BVWcpUx%2B36EjEf1F288ZFcjBqDCMh3q7klFkKpAqhoVMYUhfQ7u0oFY8I29dxK0Quy3KZXp3UcUDZm4V%2FMywqdodZZHw1zIwK2Q%2FhwC8FDYhLhxiwvz2AaQ0iAKr%2B98kw8hn4pFFB9lmF6fxcNOMStGQ4GGSa%2BGf0SR6cdCErAegeUp9Zenl4rQbl%2BOXJFOizq86q6sGhyG4GQesDCr09%2FEBjqsAoQQ5%2FhAmcWpc4g63gll0XJxWZU%2Bg2qS1aJIAcaEGahDZk%2FhA9HGG8vqVZHlle19DyRVsJQA9wRmX7Um3UXG7bDr5pU6FBXNedhn6yn2vNhsyYW%2FqdMpLWK%2BECA6d46mGdJrIFsSkB3%2FcDShsnHgIXGaq878ZCq%2BTlztlxwjNsHfHGn6blS%2FlxnslxvXTOlHHhEdIOgJ1PWaBo285fPiUzfoZmV3zD2xnUmVgYoM1BdGWD8gOpjpRINs8yFDMSn86e9wkd3k2OD1Egag4bHDxLuEyWvIM344MVA3qHo97EDWXrPvAOm3M5s5yMnrBr5keMv7ryxKweqVC%2Fu4NMsGXoMD3ZWAkRA8%2FJ%2FpxA%2B5HBGjEp8qENX%2Bo2i%2BrG%2FNdILNullLNj8awA3jiwJbEw%3D%3D&X-Amz-Signature=baf392cb3b319907fc65638a97da9f6db3a656f29959432a230d939777bacdbe&X-Amz-SignedHeaders=host&response-content-disposition=inline")
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}
