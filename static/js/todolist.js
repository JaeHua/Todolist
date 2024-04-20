<!--    页面刷新绑定数据展示-->
window.onload = function (){
    addTodoToList()
}
// 按钮事件
function handleClick() {
    //获取输入框的值
    var title = document.getElementById('title').value;

    if (title === '') {
        alert('添加失败：Title不能为空');
        return;
    }

    var data = {
        title: title
    };
    //封装成json返回给后端
    fetch('/v1/todo', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        // 必须这种异步操作，不然不能及时显示
        .then(() => {
            showSuccessAlert();
            clearInputField();
            addTodoToList();
        })
        .catch(error => {
            console.error('添加失败:', error);
        });
}
//为了美观，成功有提示框
function showSuccessAlert() {
    const alert = document.createElement('div');
    alert.classList.add(
        'fixed',
        'top-1/2',
        'left-1/2',
        'transform',
        '-translate-x-1/2',
        '-translate-y-1/2',
        'bg-green-500',
        'text-white',
        'p-4',
        'shadow-lg',
        'rounded'
    );

    const closeButton = document.createElement('button');
    //关闭提示框的按钮
    closeButton.innerHTML = '&times;';
    closeButton.classList.add('absolute',  'right-0', 'text-green', 'text-lg', 'rounded-full',  'p-1');
    closeButton.style.width = '24px';
    closeButton.style.height = '24px';
    closeButton.style.lineHeight = '1';
    closeButton.addEventListener('click', () => {
        document.body.removeChild(alert);
    });

    const text = document.createElement('p');
    text.textContent = 'Todo added successfully';
    text.classList.add('text-lg', 'font-bold');

    alert.appendChild(closeButton);
    alert.appendChild(text);

    document.body.appendChild(alert);
    //延时三秒
    setTimeout(() => {
        document.body.removeChild(alert);
    }, 3000);

}
//清除输入框
function clearInputField() {
    document.getElementById('title').value = '';
}

//显示后端数据库的数据
function addTodoToList() {
    fetch('/v1/todo')
        .then(response => response.json())
        .then(data => {
            const todoList = document.getElementById('datalist');
            todoList.innerHTML = '';

            data.forEach(todo => {
                const tr = document.createElement('tr');

                const id = document.createElement('th');
                id.textContent = todo.ID;
                id.classList.add('text-sm', 'font-semibold', 'text-gray-900');

                const title = document.createElement('th');
                title.textContent = todo.title;
                title.classList.add('text-sm', 'font-semibold', 'text-gray-900', 'ml-4');

                const action = document.createElement('th');
                const deleteButton = document.createElement('button');
                deleteButton.textContent = 'x';
                deleteButton.classList.add(
                    'bg-red-500', 'hover:bg-red-700', 'text-white', 'font-bold', 'rounded-full'
                );
                deleteButton.style.width = '30px'; // 设置宽度为30px
                deleteButton.style.height = '30px'; // 设置高度为30px
                deleteButton.style.alignItems = 'center';
                deleteButton.style.justifyContent = 'center';
                // deleteButton.onclick = () => handleDelete(todo.ID);
                action.appendChild(deleteButton);

                tr.appendChild(id);
                tr.appendChild(title);
                tr.appendChild(action);

                todoList.appendChild(tr);
            });
        });
}
