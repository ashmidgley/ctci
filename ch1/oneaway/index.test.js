const assert = require("assert");
const index = require("./index");

describe("isOneEditAway", () => {
    it("should return true", () => {
        const cases = [
            { a: "pale", b: "pales" },
            { a: "pale", b: "bale"}
        ];

        for(var i = 0; i < cases.length; i++) {
            const result = index.isOneEditAway(cases[i].a, cases[i].b);
            assert.strictEqual(true, result);
        }
    });

    it("should return false", () => {
        const cases = [
            { a: "pale", b: "paless" },
            { a: "pale", b: "pde"},
            { a: "pale", b: "palds"},
            { a: "pale", b: "bake"}
        ];

        for(var i = 0; i < cases.length; i++) {
            const result = index.isOneEditAway(cases[i].a, cases[i].b);
            assert.strictEqual(false, result);
        }
    });
});