let index = {};

index.isOneEditAway = (a, b) => {
    aLen = a.length;
    bLen = b.length;
    if(aLen !== bLen && aLen+1 !== bLen && aLen !== bLen+1) {
        return false;
    }

    if(aLen > bLen) {
        // Remove
        return offByOne(b, a);
    }
    if(aLen < bLen) {
        // Insert
        return offByOne(a, b);
    }

    //Replace
    replacementHit = false;
    for(var i = 0; i < a.length; i++) {
        if(a[i] != b[i]) {
            if(!replacementHit) {
                replacementHit = true;
            } else {
                return false;
            }
        }
    }
    return true;
};

const sortString = (s) => {
    return s.split("").sort().join("");
};

const offByOne = (a, b) => {
    a = sortString(a);
    b = sortString(b);

    extraHit = false;
    i = j = 0;
    while(i < a.length && j < b.length) {
        if(a[i] != b[j]) {
            if(!extraHit) {
                extraHit = true;
                j++;
            } else {
                return false;
            }
        } else {
            i++;
            j++
        }
    }
    return true;
};

module.exports = index;