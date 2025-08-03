import styles from "./Main.module.css";

function Main({ children }) {
	return <section className={styles.main}>{children}</section>;
}

export default Main;
