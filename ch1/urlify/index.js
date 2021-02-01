let index = {};

index.urlify = (s, len) => {
   let builder = [];
   for(var i = 0; i < len; i++) {
       if(s[i] === " ") {
           builder.push("%20");
       } else {
           builder.push(s[i]);
       }
   }
   return builder.join("");
};

module.exports = index;