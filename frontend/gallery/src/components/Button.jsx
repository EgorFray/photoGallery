import styles from "./Button.module.css";

function Button({ children, handleClick, className = "", ...props }) {
	return (
		<button
			className={`${styles.button} ${className}`}
			onClick={handleClick}
			{...props}
		>
			{children}
		</button>
	);
}

export default Button;
