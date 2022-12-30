function keyboard(index){
    let str=''
    for (let i=1; i<=1000; i++){
        str = str + i;
    }
    return str[index-1]
}


console.log(keyboard(1111))