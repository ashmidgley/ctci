const assert = require("assert");
const index = require("./index");

describe("rotate", () => {
    it("should return expected matrix", () => {
        const cases = [
            {
                in: [[1,1,1],[2,2,2],[3,3,3]],
                want: [[3,2,1],[3,2,1],[3,2,1]]
            },
            {
                in: [[1, 1], [2, 2], [3, 3]],
                want: [[3, 2, 1], [3, 2, 1]],
            },
            {
                in: [],
                want: [],
            },
        ];

        for(var i = 0; i < cases.length; i++) {
            const result = index.rotate(cases[i].in);
            assert.deepStrictEqual(cases[i].want, result);
        }
    });
});