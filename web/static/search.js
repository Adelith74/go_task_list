let search = document.getElementById("search")

    search.addEventListener("input", () => {
        let titles = document.getElementById("taskList").getElementsByClassName('title')
        Array.from(titles).forEach(element => {
            if (!element.textContent.includes(event.target.value)){
                element.parentElement.parentElement.style.display = "none"
            } else {
                element.parentElement.parentElement.style.display = "block"
            }
        });           
    })


