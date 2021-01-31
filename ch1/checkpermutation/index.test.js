const assert = require('assert');
const index = require('./index');

describe("isPermutation", () => {
    it("should return true", () => {
        const cases = [
            { a: "abc", b: "cab" },
            { a: "AbC", b: "CAb" },
            { a: "This is a permutation", b: "a permutation This is" }
        ];

        for (var i = 0; i < cases.length; i++) {
            const result = index.isPermutation(cases[i].a, cases[i].b);
            assert.strictEqual(true, result);
        };
    });

    it("should return false", () => {
        const cases = [
            { a: "abc", b: "ddd" },
            { a: "ABC", b: "abc" },
            { a: "This is a permutation", b: "a permutation This id" }
        ];

        for (var i = 0; i < cases.length; i++) {
            const result = index.isPermutation(cases[i].a, cases[i].b);
            assert.strictEqual(false, result);
        };
    });
});