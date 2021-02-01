const assert = require("assert");
const index = require("./index");

describe("compress", () => {
    it("should return expected string", () => {
        const cases = [
            { in: "aabcccccaaa", want: "a2b1c5a3" },
            { in: "aaBcccccAAa", want: "a2B1c5A2a1" },
            { in: "abcd", want: "abcd" },
            { in: "aabcccccaad", want: "a2b1c5a2d1" },
            { in: "", want: "" }
        ];
        
        for(var i = 0; i < cases.length; i++) {
            const result = index.compress(cases[i].in);
            assert.strictEqual(cases[i].want, result);
        }
    });
});