import { motion } from 'framer-motion';
interface ButtonProps {
  onClick?: (e) => void;
  className?: string;
  disabled?: boolean;
  visible?: boolean;
  padding?: "medium" | "small" | "large";
  type?: "default" | "green" | "red"
  label: string;
}
export default function Button (props: ButtonProps) {
  let padding = "";
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
  switch(props.type) {
    case "green":
      color = "bg-gradient-to-r from-green-400 to-green-600 border-green-600 shadow-green-600";
      break;
    case "red":
      color = "bg-gradient-to-r from-red-400 to-red-600 border-red-600 shadow-red-600";
      break;
    default:
      color = "bg-gradient-to-r from-blue-500 to-blue-600 border-blue-500 shadow-blue-600";
  }
  const base = "transition-colors transition-shadow duration-200 border text-slate-200 hover:text-white";
  const enabled = "hover:bg-blue-700 text-white font-bold rounded shadow-sm active:shadow-none";
  const disabled = "shadow-none bg-gradient-to-b from-gray-900 to-gray-600 bg-gradient-to-r bg-gray-500 font-bold rounded cursor-not-allowed"
  const renderClass = `${base} ${color} ${props.disabled ? disabled : enabled} ${padding} ${props.className ? props.className : ""}`;
  const clickHandler = (e) => {
    if (props.onClick && !props.disabled) {
      props.onClick(e)
    }
  }
  return (
    <motion.button
      className={`${renderClass}`}
      whileTap={!props.disabled ? { scale: 0.95 } : {}}
      onClick={clickHandler}
    >
      <span className="font-extrabold">
          {props.label || "button"}
      </span>
    </motion.button>
  )
}
