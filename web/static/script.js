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
    localStorage.setItem("tasks", JSON.stringify(data))
    let elem = document.getElementById("taskList")
    data.forEach(element => {
        let div = formTaskDiv(element)
        setTaskDivStyle(div)
        elem.appendChild(div)
    });
})
.catch(error => {
    console.error('There was a problem with your fetch operation:', error);
});

function setTaskDivStyle(div){
    div.style.margin = "20px"
    div.style.marginLeft = "50px"
    div.style.marginRight = "50px"
    div.style.boxSizing = "border-box"
    div.style.padding = "10px"
    div.style.borderRadius = "20px"
    div.style.width = "70%"
}

function formTaskDiv(element){
    var div = document.createElement("div")
    div.style.backgroundColor = "rgb(177, 146, 255)"
    var t = document.createElement("p")
    var ttdiv = document.createElement("div")
    var p = document.createElement("input")
    p.setAttribute("type", "checkbox")
    p.checked = element.status
    p.style.marginLeft = "10px"
    p.style.marginTop = "10px"
    p.addEventListener('change', function(event) {
        var parentDiv = this.parentNode;
        var el = parentDiv.querySelector('p');
        if (event.currentTarget.checked) {
            el.style.textDecoration = "line-through"
          } else {
            el.style.textDecoration = "none"
          }
    })
    ttdiv.appendChild(p)
    t.textContent = element.title
    t.style.fontFamily = "Montserrat"
    t.style.fontSize = "50px"
    t.style.padding = "5px"
    t.style.paddingLeft = "30px"
    t.style.margin = "0"
    if (element.status) {
        t.style.textDecoration = "line-through"
    }
    ttdiv.appendChild(t)
    ttdiv.style.background = "rgb(147, 124, 205)"
    ttdiv.style.borderRadius = "13px"
    ttdiv.style.margin = "3px"
    var trash = document.createElement("img")
    trash.src = "/static/trash.ico"
    trash.style.width = "32px"
    trash.style.height = "32px"
    var button = formDeleteButton()
    button.appendChild(trash)
    ttdiv.appendChild(button)
    ttdiv.style.display = "flex"
    ttdiv.style.flexDirection = "row"
    ttdiv.style.alignItems = "flex-start"
    div.appendChild(ttdiv)
    div.appendChild(formDateP(element))
    p = document.createElement("p")
    p.style.color = "black"
    p.style.fontFamily = "Montserrat"
    p.textContent = "Description: " + element.description
    div.appendChild(p)
    return div
}

function formDateP(element){
    var p = document.createElement("p")
    p.style.color = "black"
    p.style.fontFamily = "Montserrat"
    p.textContent = new Date(element.created_at).toUTCString()
    return p
}

function formTitleDiv() {

}

function formDeleteButton(){
    var button = document.createElement("button")
    button.style.backgroundColor = "rgba(0,0,0,0)"
    button.style.border = "0"
    button.style.marginLeft = "auto"
    button.style.cursor = "pointer"
    button.onclick = () => {
        window.location.href = "http://google.com";
    }
    return button
}