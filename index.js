const util = require('util');
const exec = util.promisify(require('child_process').exec);


const people = [
    {
        name: "boy.jpg",
        filter: "sepiana",
        contenttype: "image/jpg"
    },
    {
        name: "anne.jpg",
        filter: "charm",
        contenttype: "image/jpg"
    }
];

// console.log(`'${JSON.stringify(people)}'`)

async function printName() {
    return exec(`./image-filter -images '${JSON.stringify(people)}'`);
}

printName().then(data => {
    const { stdout, stderr } = data;
    console.log(stdout);
    // console.log(stderr);
});
