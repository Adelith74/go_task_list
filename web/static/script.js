let promise = fetch("http://localhost:8080/get_tasks", {
    headers: {
        'Access-Control-Allow-Origin': '*', 
        'Access-Control-Allow-Headers': 'Content-Type',
        'Access-Control-Allow-Methods': 'GET, POST, PUT'
    }
})
.then(response => {
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
})
.then(data => {
    console.log(data)
    let elem = document.getElementById("taskList")
    data.forEach(element => {
        let div = formTaskDiv(element)
        elem.appendChild(div)
    });
})
.catch(error => {
    console.error('There was a problem with your fetch operation:', error);
});

function formTaskDiv(element){
    var div = document.createElement("div")
    div.style.backgroundColor = "rgb(177, 146, 255)"
    var p = document.createElement("p")
    p.textContent = element.title
    p.style.color = "black"
    p.style.fontFamily = "Montserrat"
    div.appendChild(p)
    p = document.createElement("p")
    p.style.color = "black"
    p.style.fontFamily = "Montserrat"
    p.textContent = element.created_at
    div.appendChild(p)
    p = document.createElement("p")
    p.style.color = "black"
    p.style.fontFamily = "Montserrat"
    p.textContent = element.description
    div.appendChild(p)
    p = document.createElement("input")
    p.setAttribute("type", "checkbox")
    p.checked = element.status
    div.appendChild(p)
    return div
}