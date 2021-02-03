let index = {};

index.isPermutation = (a, b) => {
    if(a.length != b.length) {
        return false;
    }
    return sortString(a) === sortString(b);
}

const sortString = (s) => {
    return s.split("").sort().join("");
}

module.exports = index;