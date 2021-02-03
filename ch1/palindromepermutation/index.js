let index = {};

index.isPermutation = (s) => { 
    s = s.split(" ").join("").toLowerCase();
    
    occurrences = {};
    for(var i = 0; i < s.length; i++) {
        const c = s[i];
        if(!occurrences[c]) {
            occurrences[c] = 1;
        } else {
            occurrences[c]++;
        }
    }

    let singleOddHit = false;
    for(const value of Object.values(occurrences)) {
        if(value !== 2) {
            if(!singleOddHit) {
                singleOddHit = true;
            } else {
                return false;
            }
        }
    }
    return true;
};

module.exports = index;