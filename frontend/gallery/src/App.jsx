import { PostsProvider } from "./context/PostsContext";
import PostDetail from "./pages/PostDetail";
import { Routes, Route } from "react-router-dom";
import MainPage from "./pages/MainPage";

function App() {
	return (
		<div>
			<PostsProvider>
				<Routes>
					<Route path="/" element={<MainPage />} />
					<Route path="posts/:id" element={<PostDetail />} />
				</Routes>
			</PostsProvider>
		</div>
	);
}

export default App;
