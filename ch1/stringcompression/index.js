let index = {};

index.compress = (s) => {
    if(s.length === 0) {
        return s;
    }

    let builder = [];
    let count = 0;
    for(var i = 0; i < s.length-1; i++) {
        if(s[i] == s[i+1]) {
            count++;
        } else {
            builder.push(`${s[i]}${count+1}`);
            count = 0;
        }
    }
    builder.push(`${s[s.length-1]}${count+1}`);

    const result = builder.join("");
    return s.length < result.length ? s : result;
};

module.exports = index;