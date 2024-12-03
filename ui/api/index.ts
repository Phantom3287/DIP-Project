const backend = "http://localhost:8080";

export const uploadImage = async (file: File) => {
  const formData = new FormData();
  formData.append("image", file);

  const response = await fetch(`${backend}/upload`, {
    method: "POST",
    body: formData,
  });

  return response;
};

export const fast = async (threshold: number) => {
  const response = await fetch(`${backend}/fast?threshold=${threshold}`, {
    method: "GET",
  });
  if (response.status === 200) {
    console.log("FAST corner detection successful");
    const res = await response.json();
    return `${backend}/${res.path}`;
  } else {
    return null;
  }
};

export const harris = async (threshold: number) => {
  const response = await fetch(`${backend}/harris?threshold=${threshold}`, {
    method: "GET",
  });
  if (response.status === 200) {
    console.log("Harris corner detection successful");
    const res = await response.json();
    return `${backend}/${res.path}`;
  } else {
    return null;
  }
};

export const shiTomashi = async (threshold: number) => {
  const response = await fetch(`${backend}/shi-tomashi?threshold=${threshold}`, {
    method: "GET",
  });
  if (response.status === 200) {
    console.log("Shi Tomashi corner detection successful");
    const res = await response.json();
    return `${backend}/${res.path}`;
  } else {
    return null;
  }
};
