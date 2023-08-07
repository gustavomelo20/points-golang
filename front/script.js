document.addEventListener("DOMContentLoaded", function () {
   
    function fetchData() {
        const url = "http://localhost:8080/extrato?documento=48905886809";

        fetch(url, {
            headers: {
                "Content-Type": "application/json"
            }
        })
            .then((response) => response.json())
            .then((data) => displayData(data))
            .catch((error) => console.error("Erro ao buscar dados da API:", error));
    }


    function displayData(data) {

        const saldoElement = document.getElementById("saldo");
        saldoElement.textContent = `R$ ${data.saldo}`;


        const extratoElement = document.getElementById("extrato");
        extratoElement.innerHTML = ""; 

        data.Extrato.forEach((item) => {
            const listItem = document.createElement("tr");
            listItem.innerHTML = `
                <td> R$ ${item.valor}</td>
                <td> ${item.tipo}</td>
                <td> ${item.documento}</td>
            `;
            extratoElement.appendChild(listItem);
        });
    }

  
    fetchData();
});
