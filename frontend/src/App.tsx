import { CircleCheckBig, Trash2, Undo } from "lucide-react";
import { useState } from "react";
import { createTask } from "./api/taskApi";

function App() {
  const [task, setTask] = useState("");
  const todo = ["wash plates", "clean room", "eat potatoes", "call zuzu"];

  const done = ["buy car", "enter house"];

  return (
    <div className=" bg-neutral-50 flex flex-col gap-8 items-center p-10 min-h-screen w-full">
      <h1 className="text-2xl lg:text-5xl font-bold">TO-DO APP</h1>
      <form
        autoComplete="off"
        onSubmit={(e) => {
          createTask(e, task);
          setTask("");
        }}
        className=" flex justify-center h-10 items-center"
      >
        <input
          onChange={(e) => {
            setTask(e.target.value);
          }}
          type="text"
          name="task"
          value={task}
          placeholder="Add task ..."
          className="w-96 h-full border-black pl-2 border-2"
        />
        <button
          type="submit"
          className="border-2 hover:bg-neutral-400 bg-neutral-200 border-black h-full w-28"
        >
          Add Task
        </button>
      </form>
      <h2 className=" text-xl lg:text-3xl font-bold">To-Do List</h2>
      <ul className=" flex flex-col gap-2">
        {todo.map((item, index) => (
          <div className=" flex justify-center items-center gap-4">
            <li
              key={index}
              className="bg-neutral-100 w-full break-words pr-56 pl-3 py-3"
            >
              {item}
            </li>
            <button className=" hover:bg-neutral-100 p-3">
              <CircleCheckBig />
            </button>
            <button className="hover:bg-neutral-100 p-3">
              <Trash2 />
            </button>
          </div>
        ))}
      </ul>
      <h2 className=" text-xl lg:text-3xl font-bold">Done</h2>
      <ul className=" flex flex-col gap-2">
        {done.map((item, index) => (
          <div className=" flex justify-center items-center gap-4">
            <li
              key={index}
              className="bg-neutral-300 w-full break-words pr-56 line-through pl-3 py-3"
            >
              {item}
            </li>
            <button className=" hover:bg-neutral-300 p-3">
              <Undo />
            </button>
            <button className="hover:bg-neutral-300 p-3">
              <Trash2 />
            </button>
          </div>
        ))}
      </ul>
    </div>
  );
}

export default App;
