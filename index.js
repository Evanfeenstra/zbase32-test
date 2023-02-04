import zbase32 from "zbase32";

function main() {
  let arr = Buffer.from([1, 2, 3, 4, 5, 6, 7, 8]);
  const res = zbase32.encode(arr);
  console.log(res);
  // ynyodyonocbae

  let res2 = zbase32.encode(new TextEncoder("utf-8").encode("hello"));
  console.log(res2);
  // pb1sa5dx
}

main();
