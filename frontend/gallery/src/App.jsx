import { PostsProvider } from "./context/PostsContext";
import PostDetail from "./pages/PostDetail";
import { Routes, Route } from "react-router-dom";
import MainPage from "./pages/MainPage";
import CreatePost from "./pages/CreatePost";

function App() {
	return (
		<div>
			<PostsProvider>
				<Routes>
					<Route index element={<MainPage />} />
					<Route path="/create" element={<CreatePost />} />
					<Route path="posts/:id" element={<PostDetail />} />
				</Routes>
			</PostsProvider>
		</div>
	);
}

export default App;
