const assert = require("assert");
const index = require("./index");

describe("setZeroes", () => {
    it("should return expected matrix", () => {
        const cases = [
            {
                in: [[1, 0, 1], [1, 1, 1], [1, 1, 1]],
                want: [[0, 0, 0], [1, 0, 1], [1, 0, 1]],
            },
            {
                in: [[1, 0, 1, 0], [1, 1, 1, 1]],
                want: [[0, 0, 0, 0], [1, 0, 1, 0]],
            },
            {
                in: [],
                want: [],
            },
        ];

        for(var i = 0; i < cases.length; i++) {
            const result = index.setZeroes(cases[i].in);
            assert.deepStrictEqual(cases[i].want, result);
        }
    });
});