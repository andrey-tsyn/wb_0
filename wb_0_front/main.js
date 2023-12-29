function getOrder() {
    const url = document.getElementById("url").value
    const orderId = document.getElementById("orderId").value

    let content = document.getElementById("result")

    fetch(url + "?id=" + orderId)
        .then(response => {
            return response.text()
        })
        .then(data => {
            try {
                json = JSON.parse(data)
                content.innerHTML = JSON.stringify(json, null, "\t")
            } catch (e) {
                content.innerHTML = data
            }
        })
}