let index = {};

index.rotate = (m) => {
    if(m.length === 0) {
        return m;
    }

    result = [];
    for(var i = 0; i < m[0].length; i++) {
        row = [];
        for(var j = m.length-1; j >= 0; j--) {
            row.push(m[j][i]);
        }
        result.push(row);
    }

    return result;
};

module.exports = index;