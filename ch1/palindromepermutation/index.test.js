const assert = require("assert");
const index = require("./index");

describe("isPermutation", () => {
    it("should return true", () => {
        const cases = ["Tact Coa", "Tact Coao"];
        for(var i = 0; i < cases.length; i++) {
            const result = index.isPermutation(cases[i]);
            assert.strictEqual(true, result);
        }
    });

    it("should return false", () => {
        const cases = ["Tact Cod", "Tact Coaa"];
        for(var i = 0; i < cases.length; i++) {
            const result = index.isPermutation(cases[i]);
            assert.strictEqual(false, result);
        }
    });
});