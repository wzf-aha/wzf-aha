GIN测试项目
--
/**
中间件的流转
	gin提供了两个函数Abort()和Next()，配合着return关键字用来跳转或者终止存在着业务逻辑关系的中间件
	abort()就是终止该中间件的流程，如果不return的话会继续执行后面的逻辑，但不再执行其他的中间件。next()跳过当前的中间件，执行下一个中间件，待下一个中间件执行完后再回到当前next位置，直接后面的逻辑
 */