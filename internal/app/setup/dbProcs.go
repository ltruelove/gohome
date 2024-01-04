package setup

const CreateProcs string = `DROP PROCEDURE IF EXISTS public.deletenode(bigint);

CREATE OR REPLACE PROCEDURE public.deletenode(
	IN id bigint DEFAULT 0)
LANGUAGE 'sql'
AS $BODY$
Delete From viewnodesensordata Where nodeid = id;
Delete From viewnodeswitchdata Where nodeid = id;
Delete From templog WHERE nodesensorlogid IN 
	(Select id From nodesensorlog Where nodeid = id);
Delete From magneticlog WHERE nodesensorlogid IN 
	(Select id From nodesensorlog Where nodeid = id);
Delete From moisturelog WHERE nodesensorlogid IN 
	(Select id From nodesensorlog Where nodeid = id);
Delete From resistorlog WHERE nodesensorlogid IN 
	(Select id From nodesensorlog Where nodeid = id);
Delete From nodesensorlog Where nodeid = id;
Delete From nodeswitch Where nodeid = id;
Delete From nodesensor Where nodeid = id;
Delete From controlpointnodes Where nodeid = id;
Delete From node Where id = id;
$BODY$;
`
