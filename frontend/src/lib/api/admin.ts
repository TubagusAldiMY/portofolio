import { apiFetch } from './client';
import type {
	AdminResource,
	AdminResourceMap,
	ApiDataResponse,
	ApiMessageResponse,
	ContactMessage,
	UploadResponse
} from './types';

export async function listMessages(): Promise<ContactMessage[]> {
	const response = await apiFetch<ApiDataResponse<ContactMessage[]>>('/api/admin/messages');
	return response.data;
}

export async function createResource<T extends AdminResource>(
	resource: T,
	payload: Omit<AdminResourceMap[T], 'id'>
): Promise<AdminResourceMap[T]> {
	const response = await apiFetch<ApiDataResponse<AdminResourceMap[T]>>(`/api/admin/${resource}`, {
		body: payload,
		method: 'POST'
	});

	return response.data;
}

export async function updateResource<T extends AdminResource>(
	resource: T,
	id: number,
	payload: Partial<Omit<AdminResourceMap[T], 'id'>>
): Promise<AdminResourceMap[T]> {
	const response = await apiFetch<ApiDataResponse<AdminResourceMap[T]>>(
		`/api/admin/${resource}/${id}`,
		{
			body: payload,
			method: 'PUT'
		}
	);

	return response.data;
}

export async function deleteResource(
	resource: AdminResource,
	id: number
): Promise<ApiMessageResponse> {
	return apiFetch<ApiMessageResponse>(`/api/admin/${resource}/${id}`, {
		method: 'DELETE'
	});
}

export async function uploadFile(file: File): Promise<UploadResponse> {
	const formData = new FormData();
	formData.set('file', file);

	return apiFetch<UploadResponse>('/api/admin/upload', {
		body: formData,
		method: 'POST'
	});
}
