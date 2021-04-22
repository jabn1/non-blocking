var dim = 4
function insert(){
  const table = document.getElementById('t-images')

  for(let i = 0; i < dim; i++){
      let row = ''
      for(let j = 0; j < dim; j++){
      row += `<td><img id="i-${i}-${j}" ></td>`
      }
      table.innerHTML += `<tr>${row}</tr>`
  }
}
insert()
function loadSync(){
  
}
function loadAsync(){
  
  for(let i = 0; i < dim; i++){
    let row = ''
    for(let j = 0; j < dim; j++){
      let xhr = new XMLHttpRequest()
      xhr.open('GET','http://localhost:5000/api/images/test')
      let im = document.getElementById(`i-${i}-${j}`)
      xhr.send()
      xhr.onload = function(){
        im.src = xhr.responseText
      }
    }
}
}