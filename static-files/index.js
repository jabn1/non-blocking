var dim = 4
var imageName = 'ky'

function insert(){
  const table = document.getElementById('t-images')

  for(let i = 0; i < dim; i++){
      let row = ''
      for(let j = 0; j < dim; j++){
      row += `<td><img style="max-width: 300px; max-height: 400px;" id="i-${i}-${j}" ></td>`
      }
      table.innerHTML += `<tr>${row}</tr>`
  }
}
insert()
function loadSync(){
  for(let i = 0; i < dim; i++){
    for(let j = 0; j < dim; j++){
      let xhr = new XMLHttpRequest()
      xhr.open('GET','http://localhost:5000/api/images/' + imageName,false)
      let im = document.getElementById(`i-${i}-${j}`)
      xhr.send()
      im.src = xhr.responseText
    }
  }
}
function loadAsync(){ 
  for(let i = 0; i < dim; i++){
    for(let j = 0; j < dim; j++){
      let xhr = new XMLHttpRequest()
      xhr.open('GET','http://localhost:5000/api/images/' + imageName)
      let im = document.getElementById(`i-${i}-${j}`)
      xhr.send()
      xhr.onload = function(){
        im.src = xhr.responseText
      }
    }
  }
}

function loadStatic(){ 
  for(let i = 0; i < dim; i++){
    for(let j = 0; j < dim; j++){     
      document.getElementById(`i-${i}-${j}`).src = 'http://localhost:5000/images/' + imageName + '.jpg'     
    }
  }
}

function clearImages(){
  for(let i = 0; i < dim; i++){
    for(let j = 0; j < dim; j++){
      document.getElementById(`i-${i}-${j}`).src=''
    }
  }
}
