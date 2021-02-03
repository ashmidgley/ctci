let index = {};

index.isOneEditAway = (a, b) => {
    aLen = a.length;
    bLen = b.length;
    
    if(aLen+1 === bLen)
        return isInsertion(a, b);
    if(aLen === bLen+1)
        return isInsertion(b, a);
    if(aLen === bLen)
        return isReplacement(a, b);
    return false;
};

const isInsertion = (a, b) => {
    extraHit = false;
    i = j = 0;
    while(i < a.length && j < b.length) {
        if(a[i] != b[j]) {
            if(extraHit) return false;
            extraHit = true;
            j++;
        } else {
            i++;
            j++
        }
    }
    return true;
};

const isReplacement = (a, b) => {
    replacementHit = false;
    for(var i = 0; i < a.length; i++) {
        if(a[i] != b[i]) {
            if(replacementHit) return false;
            replacementHit = true;
        }
    }
    return true;
};

module.exports = index;