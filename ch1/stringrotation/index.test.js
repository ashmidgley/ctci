const assert = require("assert");
const index = require("./index");

describe("isRotation", () => {
    it("should return true", () => {
        const cases = [
            { s1: "waterbottle", s2: "erbottlewat" }
        ];

        for(var i = 0; i < cases.length; i++) {
            const result = index.isRotation(cases[i].s1, cases[i].s2);
            assert.strictEqual(true, result)
        }
    });

    it("should return false", () => {
        const cases = [
            { s1: "waterbottle", s2: "erbottleduh" },
            { s1: "waterbottle", s2: "erbottle" },
            { s1: "waterbot", s2: "erbottlewat" },
            { s1: "", s2: "" },
        ];

        for(var i = 0; i < cases.length; i++) {
            const result = index.isRotation(cases[i].s1, cases[i].s2);
            assert.strictEqual(false, result)
        }
    });
});
	