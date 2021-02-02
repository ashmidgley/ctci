let index = {};

index.isRotation = (s1, s2) => {
    if(s1.length === 0 || s1.length != s2.length) {
        return false;
    }
    
    const concat = s1 + s1;
    return concat.includes(s2);
}

module.exports = index;