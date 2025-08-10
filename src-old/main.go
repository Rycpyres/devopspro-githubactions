package main

import (
	"log"
	"net/http"
)

const imageURL = "https://rycimagespixel.s3.us-east-1.amazonaws.com/578be853144861.5929aa1414698.gif?response-content-disposition=inline&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEJL%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMSJHMEUCIQChYPO%2BgaUWxu0DMNMPd6gX%2Bj0Fxg1Y88vJy0vbnDY%2F5gIgZHdrIaHNULuOcwx26zBxzzTnKXsaloZn7ZGWngPDfoQqwgMIyv%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FARAAGgw5MDgwMjc0MDk4NTEiDDgJ4ODFzlZYiITU%2FyqWA6IrglmUNT7ASzvSh%2FAKmAqM8pbCwxod2zjXH4N8nsq8sYbGgnNoB1GqzNtuRcAdhPPMzkSBMpxntAAmoVJLiqiyMpru8AZgPXGoraCNWdyJeU89%2FByWXn3i2KlWY7avCrjvmEAvokx4siml1zJg%2B6VcDJVfgIDrdki2xTd6GhirQkISV6y1UZAYyeFIIJ%2BvEy9JPPTHp5ohsGWnoaz2TUAUb9fIViDXWE2vSo%2FmonGEz%2FtwdfaeDpyhk%2BA%2FQIQOIrt%2BsMUwyBR0933E2EJRY3S7t%2FRH%2BX8jBWVmmq5t%2BS6sqEYWw3y4jPQNjedNwHma0MAQrbh0%2FLbGP%2BEaCDLEVuyAeDgGK9uJh%2FhFhEPLbYc8FK18%2BRW1ri035DrXji0YRWGNLHta%2Bl9Pey1iqSn7fadW%2BzLpghpwqiIm9Wh5xpWEieCZWKJRlmAvwe1k9qYszbb%2FtjtgaJ%2FTbulObjVUaGKkmTmyiBOtA961t1kmBpnmvqDgpgdECIxQVB4TV66tOorCowWql0dnomCIrrTVQ6svV0aoprgwq9PfxAY63gLiegXiniRGjR%2FUprsY2P3QCKPA2sAGf03J1rzxwXFgiAqRPQ3Xn%2BBUG%2FznbABrmNiMtOAihWAgJ3YXma3agp%2BueA3bpSuUP79WA1pT4gmP3sEEGYf%2BRCN1Hpgq5Psf%2BhLACcC8j5dpuC8IgUFKgvmAlxL%2BoI3RiF3EOM5lA9FP48wPRSj9fDtKbl3hbNdVOErAaaJX6NuPJq%2BSiNjM8YIW8ZQbv9iigGe%2BYSLQbI8i09zeOTSoMhx7LGqT3VAFidqAyP90Gd9pXfiqWzRFWu9E3uszjUb9AgRvEFQl3%2Fx50HWJAuKYE%2Bs8CMJk0iGn31MMMWhCCwQwOCWra3d38j6UZ%2FLXhPyRq3Yv9bdxEyEgmWvSDU1oihoU7Ak6EaiPrfcksnOKj8kSOiYLLNdbQ41byFhXB6G9kYzrhHcNgOAh3nuOVDyAc%2BfVbrS1G%2BGsMUSzRcEIcANdJwRoXzzzUQ%3D%3D&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=ASIA5G2VGWW5TY73WYY3%2F20250810%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20250810T010728Z&X-Amz-Expires=3600&X-Amz-SignedHeaders=host&X-Amz-Signature=240f74d2be05cb155f639a6a01f9d28bba300d0ecb05328662f206137652da43"

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, imageURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}
