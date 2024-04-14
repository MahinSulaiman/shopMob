
//function to search 
async function searchProducts() {
    const searchQuery = document.getElementById('searchInput').value.trim();

    if (searchQuery === '') {
        alert('Please enter a search query');
        return;
    }

    try {
        const response = await fetch(`http://localhost:8080/searchPdt?query=${encodeURIComponent(searchQuery)}`);
        if (!response.ok) {
            throw new Error('Failed to search products');
        }
        const products = await response.json();
        let searchHtml = '';
products.forEach(product => {
searchHtml += `
<div class="product">
<h2>${product.name}</h2>
<p><strong>Description:</strong> ${product.speccs}</p>
<img src="${product.imageUrl}" alt="${product.name}">
</div>
`;
});
        document.getElementById('product-list').innerHTML = searchHtml;
    } catch (error) {
        console.error('Error:', error);
    }
    
}

// Function to fetch and display product data
async function fetchProductData() {
    try {
        const response = await fetch('http://localhost:8080/listPdts');
        const data = await response.json();
        
      
        const productListContainer = document.getElementById('product-list');

        // Loop  products 
        data.forEach(product => {
            // Create a container div for each product
            const productContainer = document.createElement('div');
            productContainer.classList.add('product-item');
            productContainer.id = `product-${product.id}`; 

            // Create elements 
            const productName = document.createElement('h2');
            productName.textContent = product.name;

            const productPrice = document.createElement('p');
            productPrice.textContent = `Price: $${product.price}`;

            const productImage = document.createElement('img');
            productImage.src = product.imageUrl;
            productImage.alt = product.name;

            // Create  button
            const editButton = document.createElement('button');
            editButton.textContent = 'Edit';
            editButton.onclick = function() {
                editProduct(product.id);
            };

            // Create delete button
            const deleteButton = document.createElement('button');
            deleteButton.textContent = 'Delete';
            deleteButton.onclick = function() {
                deleteProduct(product.id);
            };

            // Create  view button
            const viewButton = document.createElement('button');
            viewButton.textContent = 'View';
            viewButton.onclick = function() {
                viewProduct(product.id);
            };

            // Append  to the product container
            productContainer.appendChild(productName);
            productContainer.appendChild(productPrice);
            productContainer.appendChild(productImage);
            productContainer.appendChild(editButton);
            productContainer.appendChild(deleteButton);
            productContainer.appendChild(viewButton);

            // Append the product  to the product list 
            productListContainer.appendChild(productContainer);
        });
    } catch (error) {
        console.error('Error fetching product data:', error);
    }
}

// Fn to view product details
function viewProduct(productId) {
    // pass id to next page
    window.location.href = `product-details.html?id=${productId}`;
}

// Function to edit 
function editProduct(productId) {
    // pass id to list page
    window.location.href = `edit-product.html?id=${productId}`;
}

// Fn to delete product
async function deleteProduct(productId) {
    try {
        const response = await fetch(`http://localhost:8080/deletePdt?id=${productId}`, {
            method: 'DELETE'
        });
        
        
        if (response.status === 200) {
            
            window.alert("deleted")
           
        }
    } catch (error) {
        console.error('Error deleting product:', error);
    }
    window.alert("deleted")
    window.location.reload();
}


window.onload = fetchProductData;
