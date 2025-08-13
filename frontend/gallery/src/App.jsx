import { PostsProvider } from "./context/PostsContext";
import PostDetail from "./pages/PostDetail";
import { Routes, Route } from "react-router-dom";
import MainPage from "./pages/MainPage";
import { PostDetailProvider } from "./context/PostDetailContext";

function App() {
	return (
		<div>
			<PostsProvider>
				<PostDetailProvider>
					<Routes>
						<Route path="/" element={<MainPage />} />
						<Route path="posts/:id" element={<PostDetail />} />
					</Routes>
				</PostDetailProvider>
			</PostsProvider>
		</div>
	);
}

export default App;
