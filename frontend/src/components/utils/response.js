export function handleSuccess(response) {
  const { data, message } = response.data;
  return { success: true, data, message };
}

export function handleError(error) {
  const res = error.response?.data;

  return {
    success: false,
    error: res?.error || "Unknown Error",
    details: res?.details || null,
    status: error.response?.status || 500,
  };
}