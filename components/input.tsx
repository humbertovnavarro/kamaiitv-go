import { motion } from 'framer-motion';
import HighlightOffIcon from '@mui/icons-material/HighlightOff';
import { useState } from 'react';
interface InputProps {
  onClick?: (e) => void;
  className?: string;
  disabled?: boolean;
  visible?: boolean;
  padding?: "medium" | "small" | "large";
  color?: "default" | "green" | "red"
  label: string;
}
export function TextInput(props: InputProps) {
  let padding = "";
  const [value, setValue] = useState("");
  switch(props.padding) {
    case "small":
      padding = "py-1 px-2";
      break;
    case "medium":
      padding = "py-2 px-4";
      break;
    case "large":
      padding = "py-3 px-6";
      break;
    default:
      padding = "py-1 px-2";
      break;
  }
  let color = "";
  switch(props.color) {
    case "green":
      color = "border-green-600 shadow-green-600 shadow-md";
      break;
    case "red":
      color = "border-red-600 shadow-red-600 shadow-md";
      break;
    default:
      color = "focus:shadow-blue-600";
  }
  const base = "transition-all duration-200 border bg-stone-800 text-white outline-none focus:shadow-md";
  const enabled = "text-white font-bold rounded";
  const disabled = "shadow-none bg-gray-600 font-bold rounded-none cursor-not-allowed select-none"
  const renderClass = `${base} ${color} ${props.disabled ? disabled : enabled} ${padding} ${props.className ? props.className : ""}`;
  return (
    <span className="relative">
      <input
        className={`${renderClass} relative pl-2 pr-10`}
        disabled={props.disabled}
        value={value}
        onChange={(e) => setValue(e.target.value)}
        type={props.type === "password" ? "password" : "text"}
      >
      </input>
      <HighlightOffIcon
        onClick={() => {setValue("")}}
        className='absolute right-3 bottom-[-1px] hover:text-white transition-all duration-200 ease-in-out active:translate-y-[1px]'
      />
    </span>
  )
}
