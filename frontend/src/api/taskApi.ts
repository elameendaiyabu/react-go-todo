import axios from "axios";
import React from "react";

export async function createTask(e: React.SyntheticEvent, task: string) {
  e.preventDefault();

  try {
    const response = await axios.post(
      "http://localhost:8080",
      { task },
      {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
      },
    );
    console.log(response);
  } catch (e) {
    console.log(e);
  }
}

export async function getTasks() {
  try {
    const { data } = await axios.get("http://localhost:8080");
    //    console.log(data);
    return data;
  } catch (e) {
    console.log(e);
  }
}
