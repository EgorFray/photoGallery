import styles from "./Button.module.css";
import classNames from "classnames";

function Button({ children, handleClick, className, ...props }) {
	return (
		<button
			className={classNames(styles.button, className)}
			onClick={handleClick}
			{...props}
		>
			{children}
		</button>
	);
}

export default Button;
