import { useState } from "react";
import CreatePostForm from "../components/CreatePostForm";
import Footer from "../components/Footer";
import Header from "../components/Header";
import List from "../components/List";
import Main from "../components/Main";
import NavBar from "../components/NavBar";

function MainPage() {
	const [isOpen, setIsOpen] = useState(false);

	function toggleForm() {
		setIsOpen(!isOpen);
	}

	return (
		<div>
			<NavBar />
			<Header />
			<Main>
				<List />
			</Main>
			<Footer />
		</div>
	);
}

export default MainPage;
