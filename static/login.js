async function getSID() {
	let res;
	try {
		res = await fetch("/api/login", {
			method: "POST",
			body: JSON.stringify({
				name: document.querySelector("#name").value,
				password: document.querySelector("#password").value,
			}),
		});
	} catch (e) {
		console.error(e);
		return;
	}

	let data = await res.json();
	sessionStorage.setItem("name", data.name);
	sessionStorage.setItem("sid", data.sid);
	sessionStorage.setItem("id", data.id);

	window.location.href = "/";
}
