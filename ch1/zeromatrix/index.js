let index = {};

index.setZeroes = (m) => {
    if(m.length === 0) {
        return m;
    }

    xZeroes = new Set();
    yZeroes = new Set();
    for(var i = 0; i < m.length; i++) {
        for(var j = 0; j < m[i].length; j++) {
            if(m[i][j] === 0) {
                xZeroes.add(j);
                yZeroes.add(i);
            }
        }
    }

    for(var y = 0; y < m.length; y++) {
        xZeroes.forEach(x => {
            m[y][x] = 0;
        });
    }

    yZeroes.forEach(y => {
        for(var x = 0; x < m[0].length; x++) {
            m[y][x] = 0;
        }
    });

    return m;
}

module.exports = index;