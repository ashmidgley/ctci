const assert = require('assert');
const index = require('./index');

describe("urlify", () => {
    it("should return expected string", () => {
        const cases = [
            { in: "test test  ", len: 9, want: "test%20test" },
            { in: "Mr John Smith    ", len: 13, want: "Mr%20John%20Smith" },
            { in: "", len: 0, want: "" },
            { in: " ", len: 1, want: "%20" }
        ];

        for(var i = 0; i < cases.length; i++) {
            const result = index.urlify(cases[i].in, cases[i].len);
            assert.strictEqual(cases[i].want, result);
        }
    });
})