import { useState } from "react";
import { PostsProvider } from "./context/PostsContext";
import NavBar from "./components/NavBar";
import Header from "./components/Header";
import Search from "./components/Search";
import CreatePostForm from "./components/CreatePostForm";
import Main from "./components/Main";
import List from "./components/List";
import PostDetail from "./components/PostDetail";
import Footer from "./components/Footer";
import { Routes, Route } from "react-router-dom";

function App() {
	const [isOpen, setIsOpen] = useState(false);

	function toggleForm() {
		setIsOpen(!isOpen);
	}

	return (
		<div>
			<PostsProvider>
				<NavBar />
				<Header onOpen={toggleForm} />
				<Search />
				{isOpen && <CreatePostForm onOpen={toggleForm} />}
				<Main>
					<Routes>
						<Route path="/" element={<List />} />
						<Route path="posts/:id" element={<PostDetail />} />
					</Routes>
				</Main>
				<Footer />
			</PostsProvider>
		</div>
	);
}

export default App;
