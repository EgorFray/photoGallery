import { PostsProvider } from "./context/PostsContext";
import PostDetail from "./pages/PostDetail";
import { Routes, Route } from "react-router-dom";
import MainPage from "./pages/MainPage";
import CreatePost from "./pages/CreatePost";
import PageNotFound from "./pages/PageNotFound";
import Homepage from "./pages/Homepage";

function App() {
	return (
		<div>
			<PostsProvider>
				<Routes>
					<Route index element={<MainPage />} />
					<Route path="/homepage" element={<Homepage />} />
					<Route path="/create" element={<CreatePost />} />
					<Route path="posts/:id" element={<PostDetail />} />
					<Route path="*" element={<PageNotFound />} />
				</Routes>
			</PostsProvider>
		</div>
	);
}

export default App;
