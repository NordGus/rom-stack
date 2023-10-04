import "htmx.org";
import "./style.css"
import "./components/CounterComponent.ts"

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <div class="flex flex-row justify-center">
    <a href="https://vitejs.dev" target="_blank" class="text-indigo-500 font-medium hover:text-indigo-900">
      <img src="http://localhost:5173/vite.svg" class="h-[6rem] p-[1.5rem] hover:drop-shadow-[0_0_2rem_#646cffaa] will-change-[filter] transition-[filter] duration-300" alt="Vite logo" />
    </a>
    <a href="https://www.typescriptlang.org/" target="_blank" class="text-indigo-500 hover:text-indigo-900">
      <img src="http://localhost:5173/typescript.svg" class="h-[6rem] p-[1.5rem] hover:drop-shadow-[0_0_2rem_#3178c6aa] will-change-[filter] transition-[filter] duration-300" alt="TypeScript logo" />
    </a>
  </div>
  <div>
    <h1 class="text-[3.2rem] leading-[1.1] font-[system-ui]">Vite + TypeScript</h1>
    <div class="p-[2em]">
      <button
        type="button"
        is="rom-counter" 
        class="rounded-[8px] border-[1px] border-transparent hover:border-indigo-500 py-[0.6em] px-[1.2em] text-[1em] 
            font-medium bg-stone-900 cursor-pointer transition-[border-color] duration-[0.25s] focus:focus-visible:outline-4 
            focus:focus-visible:outline focus:focus-visible:outline[-webkit-focus-ring-color]"
      ></button>
    </div>
    <p>Click on the Vite and TypeScript logos to learn more</p>
  </div>
`
