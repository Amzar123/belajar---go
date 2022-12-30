function planetThanos(days) {
    let value = 1;
    
    for (let i = 2; i <= days; i++) {
        
        if (i % 3 === 0) {
            value = Math.floor(value/2);
        } else {
            value = value*3;
        }
    }
    return value;
}

console.log(planetThanos(50))