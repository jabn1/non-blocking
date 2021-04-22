function insert(){
  const table = document.getElementById('t-images')
  const dim = 4
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

}