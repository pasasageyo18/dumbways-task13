// VARIABLE

// var
// let
// const

// TIPE DATA

// let umur = 22; // number
// let nama = "Muhammad Nur Faza"; // string
// let havingJob = false; // boolean

// console.log(`nama saya ${nama}. Umur saya ${umur}`);
// console.log("nama saya", nama, "umur saya", umur, "tahun");
// console.log("nama saya" + nama + "umur saya" + umur + "tahun");

// FUNCTION

// function hello() {
//   let x = 5;
//   let y = 10;
//   let result = x * y;

//   console.log(result);
// }

// hello();

// function hello2(x, y) {
//   x = 5;
//   y = 2;
//   result = x * y;
//   console.log(result);
// }

function submitData() {
  let name = document.getElementById("nama").value;
  let email = document.getElementById("email").value;
  let phone = document.getElementById("noHp").value;
  let subject = document.getElementById("subject").value;
  let message = document.getElementById("message").value;

  if (name == "") {
    return alert("Nama harus diisi!");
  } else if (email == "") {
    return alert("Email harus diisi!");
  } else if (phone == "") {
    return alert("Nomor HP harus diisi!");
  } else if (subject == "") {
    return alert("Subject harus dipilih!");
  } else if (message == "") {
    return alert("Pesan harus diisi!");
  }

  let emailReceiver = "gofar222@gmail.com";

  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name}. ${message}. Silahkan kontak saya di nomor ${phone}, terima kasih.`;
  a.click();

  let emailer = {
    name,
    email,
    phone,
    subject,
    message,
  };

  console.log(emailer);

  console.log(name, email, phone, subject, message);
}
