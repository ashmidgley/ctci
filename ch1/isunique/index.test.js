const assert = require('assert');
const index = require('./index');

describe('allUnique', () => {
    it('should return true', () => {
        const cases = ["abcdefg", "世界"];
        for (var i = 0; i < cases.length; i++) {
            const result = index.allUnique(cases[i]);
            assert.strictEqual(true, result);
        };
    });

    it('should return false', () => {
        const cases = ["aaaaaaa", "世世"];
        for (var i = 0; i < cases.length; i++) {
            const result = index.allUnique(cases[i]);
            assert.strictEqual(false, result);
        };
    });
});