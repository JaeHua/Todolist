
    // 页面刷新绑定数据展示
    window.onload = function (){
    addTodoToList()
}
    //enter绑定
    // 获取输入框元素
    const inputField = document.getElementById('title');

    // 绑定按键事件监听器
    inputField.addEventListener('keydown', function(event) {
    if (event.key === 'Enter') {
    event.preventDefault();
    handleClick();
}
});
    inputField.focus()
    // 按钮事件
    function handleClick() {
    //获取输入框的值
    let title = document.getElementById('title').value;

    if (title === '') {
    alert('添加失败：Title不能为空');
    return;
}

    let data = {
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
    'inset-x-10',
    'top-0',
    'mt-2',
    'left-1/3',
    'right-1/4',
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
    text.classList.add('text-lg','text-center', 'font-bold');

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
    //更新操作
    function handleCheck(id) {
    const checkButton = document.getElementById(`checkButton-${id}`);
    const isCompleted = checkButton.classList.contains('completed');
    // 发起请求
    fetch(`/v1/todo/${id}`, {
    method: 'PUT',
    headers: {
    'Content-Type': 'application/json'
},
})

    .then(() => {
    if (isCompleted) {
    checkButton.classList.remove('completed');
} else {
    checkButton.classList.add('completed');
}
})
    .then(() => {
    addTodoToList(); // 刷新待办事项列表

})
    //及时修改返回结果，免得 window.location.reload();自主刷新很麻烦
    .then(response => response.json())
    .catch(error => {
    console.error('添加失败:', error);
});

}
    //删除操作
    function handleDelete(id){
    const deleteButton = document.getElementById(`deleteButton-${id}`);
    fetch(`/v1/todo/${id}`, {
    method: 'DELETE',
    headers: {
    'Content-Type': 'application/json'
},
})
    .then(response => response.json())
    .then(() => {
    addTodoToList(); // 刷新待办事项列表

})
    //及时修改返回结果，免得 window.location.reload();自主刷新很麻烦
    .then(response => response.json())
    .catch(error => {
    console.error('删除失败:', error);
});
}
    //显示后端数据库的数据
    function addTodoToList() {
    var idd = 0
    fetch('/v1/todo')
    .then(response => response.json())
    .then(data => {
    const todoList = document.getElementById('datalist');
    todoList.innerHTML = '';

    data.forEach(todo => {
    const tr = document.createElement('tr');

    const id = document.createElement('th');
    id.textContent = ++idd;
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
    deleteButton.id = `deleteButton-${todo.ID}`; // 设置按钮的id，用于后续操作


    //创建打钩按钮
    const checkButton = document.createElement('button');
    checkButton.textContent = '✓';
    checkButton.id = `checkButton-${todo.ID}`; // 设置按钮的id，用于后续操作
    checkButton.classList.add(
    'bg-green-500', 'hover:bg-green-700', 'text-white', 'font-bold', 'rounded-full'
    );
    checkButton.style.width = '30px';
    checkButton.style.height = '30px';
    checkButton.style.alignItems = 'center';
    checkButton.style.justifyContent = 'center';

    // 判断状态并设置按钮样式和行为
    if (todo.status) {
    checkButton.textContent = '\u27F3';
    checkButton.style.backgroundColor = 'orange';
    checkButton.classList.add('completed');
} else {
    checkButton.textContent = '✓';
    checkButton.style.backgroundColor = 'bg-green-500';
    checkButton.classList.remove('completed');
}

    //绑定事件
    checkButton.onclick = ()=> handleCheck(todo.ID)
    action.appendChild(checkButton);
    //空格
    const space = document.createTextNode(' ');
    action.appendChild(space);
    deleteButton.onclick = () => handleDelete(todo.ID);
    action.appendChild(deleteButton);

    tr.appendChild(id);
    tr.appendChild(title);
    tr.appendChild(action);

    todoList.appendChild(tr);
});
});

}
