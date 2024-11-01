import { check } from "k6";
import http from "k6/http";
import { Rate } from "k6/metrics";

export let errorRate = new Rate("errors");

export default function () {
	var url = "http://localhost:8080/api/create";

	var params = {
		headers: {
			"Content-Type": "application/json",
		},
	};

	var data = JSON.stringify({
		url: "https://www.youtube.com",
        user_id: "test_user",
	});

	check(http.post(url, data, params), {
		"status is 200": (r) => r.status == 200,
	}) || errorRate.add(1);
}
